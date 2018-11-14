import request from '@/utils/request'

export function fetchList (query) {
  return request({
    url: '/v1/courses',
    method: 'get',
    params: query
  })
}

export function fetchSectionList (courseId) {
  return request({
    url: '/v1/courses/' + courseId + '/sections',
    method: 'get'
  })
}

export function fetchCourse (id) {
  return request({
    url: '/v1/courses/' + id,
    method: 'get',
    params: { id }
  })
}

export function createCourse (data) {
  return request({
    url: '/v1/courses',
    method: 'post',
    data
  })
}

export function updateCourse (data) {
  return request({
    url: '/v1/courses/' + data.id,
    method: 'put',
    data
  })
}

export function updateSection (data) {
  return request({
    url: '/v1/courses/' + data.id,
    method: 'put',
    data
  })
}
