package service

import (
	"html/template"
	"sync"

	"github.com/1024casts/1024casts/util"

	"github.com/1024casts/1024casts/model"
	"github.com/1024casts/1024casts/repository"

	"github.com/lexkong/log"
)

type WikiService struct {
	repo *repository.WikiRepo
}

type CategoryInfo struct {
	Id        int                    `json:"id"`
	Name      string                 `json:"name"`
	WikiPages []*model.WikiPageModel `json:"wiki_pages"`
}

func NewWikiService() *WikiService {
	return &WikiService{
		repository.NewWikiRepo(),
	}
}

// 所有wiki的分类
func (srv *WikiService) GetCategoryList() ([]*model.WikiCategoryModel, error) {
	categories, err := srv.repo.GetCategoryList()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (srv *WikiService) GetWikiById(id int) (*model.WikiPageModel, error) {
	Wiki, err := srv.repo.GetWikiById(id)

	if err != nil {
		return Wiki, err
	}

	return Wiki, nil
}

func (srv *WikiService) GetWikiPageBySlug(slug string) (*model.WikiPageInfo, error) {
	wikiPageModel, err := srv.repo.GetWikiBySlug(slug)
	page := srv.trans(wikiPageModel)

	if err != nil {
		return page, err
	}

	return page, nil
}

func (srv *WikiService) GetWikiCategoryListWithPage() ([]*model.WikiCategoryModel, error) {
	infos := make([]*model.WikiCategoryModel, 0)

	categories, err := srv.repo.GetCategoryList()
	if err != nil {
		return nil, err
	}

	ids := []uint64{}
	for _, cate := range categories {
		ids = append(ids, cate.Id)
	}

	wg := sync.WaitGroup{}
	categoryList := model.WikiCategoryList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.WikiCategoryModel, len(categories)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// Improve query efficiency in parallel
	for _, c := range categories {
		wg.Add(1)
		go func(category *model.WikiCategoryModel) {
			defer wg.Done()

			pageList, err := srv.repo.GetWikiPageListByCategoryId(category.Id)
			if err != nil {
				log.Warnf("[course] get wiki page list fail from wiki repo, category_id: %d", c.Id)
				errChan <- err
				return
			}

			categoryList.Lock.Lock()
			defer categoryList.Lock.Unlock()

			for _, page := range pageList {
				category.WikiPages = append(category.WikiPages, srv.trans(page))
			}

			categoryList.IdMap[category.Id] = category
		}(c)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, err
	}

	for _, id := range ids {
		infos = append(infos, categoryList.IdMap[id])
	}

	return infos, nil
}

func (srv *WikiService) trans(page *model.WikiPageModel) *model.WikiPageInfo {
	return &model.WikiPageInfo{
		Id:         page.Id,
		CategoryId: page.CategoryId,
		Slug:       page.Slug,
		Title:      page.Title,
		Content:    template.HTML(page.Content),
		ViewCount:  page.ViewCount,
		FixCount:   page.FixCount,
		CreatedAt:  util.TimeToDateString(page.CreatedAt),
		UpdatedAt:  util.TimeToString(page.UpdatedAt),
	}
}

func (srv *WikiService) UpdateWiki(WikiMap map[string]interface{}, id int) error {
	err := srv.repo.UpdateWiki(WikiMap, id)

	if err != nil {
		return err
	}

	return nil
}
