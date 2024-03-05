import { http } from '@/utils/http/axios';

export function GetCarModelList(params?) {
  return http.request({
    url: '/business/car/model/list',
    method: 'GET',
    params,
  });
}

export function DeleteCarModel(ids: number[]) {
  return http.request({
    url: '/business/car/model/delete',
    method: 'DELETE',
    data: { ids: ids },
  });
}

export function EditCar(data) {
  return http.request(
    {
      url: '/business/car/model/edit',
      method: 'POST',
      data: data,
    },
    { isTransformResponse: false }
  );
}

export function Status(id: number, status: number) {
  return http.request({
    url: '/business/car/model/status',
    method: 'PUT',
    data: { id, status },
  });
}
