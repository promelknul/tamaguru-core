/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { AuthResp } from '../models/AuthResp';
import type { LoginReq } from '../models/LoginReq';
import type { RegisterReq } from '../models/RegisterReq';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class DefaultService {
    /**
     * @param requestBody
     * @returns AuthResp OK
     * @throws ApiError
     */
    public static postAuthRegister(
        requestBody: RegisterReq,
    ): CancelablePromise<AuthResp> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/auth/register',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * @param requestBody
     * @returns AuthResp OK
     * @throws ApiError
     */
    public static postAuthLogin(
        requestBody: LoginReq,
    ): CancelablePromise<AuthResp> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/auth/login',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
}
