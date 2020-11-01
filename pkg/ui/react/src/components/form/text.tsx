import React from "react";

import styles from "./style.module.css";

export type TextProps = {
  inpRef?:
    | string
    | ((instance: HTMLInputElement | null) => void)
    | React.RefObject<HTMLInputElement>
    | null
    | undefined;
  type?: "text" | "number" | "password" | "email" | "date";
} & React.DetailedHTMLProps<React.InputHTMLAttributes<HTMLInputElement>, HTMLInputElement>;

export const Text: React.FC<TextProps> = ({ inpRef, type = "text", className, ...rest }) => {
  return <input {...rest} ref={inpRef} type={type} className={`${styles.text} ${className}`} />;
};

export default Text;
