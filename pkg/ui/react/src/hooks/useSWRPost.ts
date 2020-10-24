import { useState, Dispatch, SetStateAction } from "react";
import useSWR, { ConfigInterface, responseInterface } from "swr";

const useSWRPost = <S>(
  endpoint: RequestInfo,
  swrOpts: ConfigInterface,
  method: string = "POST"
): [Dispatch<SetStateAction<S | undefined>>, responseInterface<any, any>] => {
  const [values, runFetch] = useState<S>();

  const swrOut = useSWR(values ? [endpoint, method, values] : null, {
    revalidateOnFocus: false,
    ...swrOpts,
  });

  return [runFetch, swrOut];
};

export default useSWRPost;
