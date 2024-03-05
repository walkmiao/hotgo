import { http } from '@/utils/http/axios';

export function GetCustomerList(params?) {
  return http.request({
    url: '/business/customer',
    method: 'GET',
    params,
  });
}

export function DeleteCustomer(ids: number[]) {
  return http.request({
    url: '/business/customer',
    method: 'DELETE',
    data: { ids: ids },
  });
}

export function EditCustomer(data) {
  return http.request(
    {
      url: '/business/customer',
      method: 'POST',
      data: data,
    },
    { isTransformResponse: false }
  );
}
export function GetCustomerOption() {
  return http.request({
    url: '/business/customer/option',
    method: 'GET',
  });
}

export function Status(id: number, status: number) {
  return http.request({
    url: '/business/customer/status',
    method: 'PUT',
    data: { id, status },
  });
}

export function GetCustomerContactList(id: number) {
  return http.request({
    url: '/business/customer/contact',
    method: 'GET',
    params: { id },
  });
}
