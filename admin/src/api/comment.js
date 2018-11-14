import request from '@/utils/request'

export function fetchList (query) {
  return request({
    url: '/v1/comments',
    method: 'get',
    params: query
  })
}
