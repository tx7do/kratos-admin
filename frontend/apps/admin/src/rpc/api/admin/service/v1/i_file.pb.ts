// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.0
//   protoc               unknown
// source: admin/service/v1/i_file.proto

/* eslint-disable */
import { Observable } from "rxjs";
import {
  type OssUploadUrlRequest,
  type OssUploadUrlResponse,
  type UploadFileRequest,
  type UploadFileResponse,
} from "../../../file/service/v1/file.pb";

/** 文件服务 */
export interface FileService {
  /** 获取对象存储（OSS）上传用的预签名链接 */
  OssUploadUrl(request: OssUploadUrlRequest): Promise<OssUploadUrlResponse>;
  /** POST方法上传文件 */
  PostUploadFile(request: Observable<UploadFileRequest>): Promise<UploadFileResponse>;
  /** PUT方法上传文件 */
  PutUploadFile(request: Observable<UploadFileRequest>): Promise<UploadFileResponse>;
}
