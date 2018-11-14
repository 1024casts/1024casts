import request from '@/utils/request'

export function fetchList (query) {
  return request({
    url: '/v1/users',
    method: 'get',
    params: query
  })
}

export function fetchArticle (id) {
  return request({
    url: '/v1/users/' + id,
    method: 'get',
    params: { id }
  })
}

export function updateStatus (id, status) {
  return request({
    url: '/v1/users/' + id + '/status',
    method: 'put',
    data: {status}
  })
}
