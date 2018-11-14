import request from '@/utils/request'

export function fetchList (query) {
  return request({
    url: '/v1/orders',
    method: 'get',
    params: query
  })
}
