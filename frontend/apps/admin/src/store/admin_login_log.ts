import { defineStore } from 'pinia';

import { defAdminLoginLogService, makeQueryString } from '#/rpc';

export const useAdminLoginLogStore = defineStore('admin_login_log', () => {
  /**
   * 查询登录日志列表
   */
  async function listAdminLoginLog(
    page: number,
    pageSize: number,
    formValues: object,
    fieldMask: null | string = null,
    orderBy: string[] = [],
    noPaging: boolean = false,
  ) {
    return await defAdminLoginLogService.ListAdminLoginLog({
      // @ts-ignore proto generated code is error.
      fieldMask,
      orderBy,
      query: makeQueryString(formValues),
      page,
      pageSize,
      noPaging,
    });
  }

  /**
   * 查询登录日志
   */
  async function getAdminLoginLog(id: number) {
    return await defAdminLoginLogService.GetAdminLoginLog({ id });
  }

  return {
    listAdminLoginLog,
    getAdminLoginLog,
  };
});

/**
 * 成功失败的颜色
 * @param success
 */
export function successToColor(success: boolean) {
  return success ? 'green' : 'red';
}

export function successToName(success: boolean, statusCode: number) {
  return success ? '成功' : `失败 (${statusCode})`;
}
