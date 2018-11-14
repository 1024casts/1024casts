import request from '@/utils/request'

export function fetchList (query) {
  return request({
    url: '/v1/plans',
    method: 'get',
    params: query
  })
}

export function fetchPlan (id) {
  return request({
    url: '/v1/plans/' + id,
    method: 'get',
    params: { id }
  })
}

export function createPlan (data) {
  return request({
    url: '/v1/plans',
    method: 'post',
    data
  })
}

export function updatePlan (data) {
  return request({
    url: '/v1/plans/' + data.id,
    method: 'put',
    data
  })
}
