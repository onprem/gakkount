import { useEffect } from "react";
import { useToasts } from "react-toast-notifications";
import { FieldError, DeepMap } from "react-hook-form";

const useFormErrors = (errors: DeepMap<Record<string, unknown>, FieldError>) => {
  const { addToast } = useToasts();

  useEffect(() => {
    const err = Object.values(errors).find((e) => e?.message);
    if (err?.message) addToast(err?.message, { appearance: "error" });
  }, [errors, addToast]);
};

export default useFormErrors;
