import React from "react";

import styles from "./style.module.css";

export type SelectProps = {
  inpRef?:
    | string
    | ((instance: HTMLSelectElement | null) => void)
    | React.RefObject<HTMLSelectElement>
    | null
    | undefined;
} & React.DetailedHTMLProps<React.InputHTMLAttributes<HTMLSelectElement>, HTMLSelectElement>;

export const Select: React.FC<SelectProps> = ({ inpRef, className, children, ...rest }) => {
  console.log({ inpRef });
  return (
    <select {...rest} ref={inpRef} className={`${styles.select} ${className}`}>
      {children}
    </select>
  );
};

export default Select;
