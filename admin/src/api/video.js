import request from '@/utils/request'

export function fetchVideoList (courseId) {
  return request({
    url: '/v1/videos/' + courseId,
    method: 'get'
  })
}
