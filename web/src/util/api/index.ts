/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable no-useless-catch */
import { Config } from '../../config';
import { joinUrl } from '../join-url';
import { getPostExpensePlanListUrl, postExpensePlanListResponse } from './generated/generated';
import { HandlerListExpensePlanRequest } from './generated/generated.schemas';

const restEndpoint = Config.REST_ENDPOINT;

async function defaultPostQuery(
  url: string,
  request: any,
  timeout = Config.DEFAULT_REST_TIMEOUT_MS
) {
  const controller = new AbortController();

  const id = setTimeout(() => {
    controller.abort('Query Timeout');
  }, timeout);
  try {
    const response = await fetch(url, {
      signal: controller.signal,
      method: 'POST',
      body: JSON.stringify(request),
    });

    // if (!response.ok) {
    //   console.log('defaultPostQuery Error', response);
    //   const data = await response.json();
    //   const e = new ApiError(data.error_code);
    //   e.error_code = data.error_code;
    //   e.error_kind = data.error_kind;
    //   e.status_code = response.status;
    //   console.error('defaultPostQuery handled Error', e);
    //   throw e;
    // }

    return { response };
  } catch (err) {
    // if (err.name === 'AbortError') {
    //   console.error('Request timed out', err);
    //   const timeoutErrorCode = 'api_timeout';
    //   const e = new ApiError(timeoutErrorCode);
    //   e.error_code = timeoutErrorCode;
    //   e.error_kind = timeoutErrorCode;
    //   e.status_code = 408;
    //   throw e;
    // }

    throw err;
  } finally {
    clearTimeout(id);
  }
}

export async function restListExpensePlan(
  request: HandlerListExpensePlanRequest,
  timeout = Config.DEFAULT_REST_TIMEOUT_MS
): Promise<Required<postExpensePlanListResponse['data']>> {
  const url = joinUrl(restEndpoint, getPostExpensePlanListUrl());
  const { response } = await defaultPostQuery(url, request, timeout);
  return response.json();
}
