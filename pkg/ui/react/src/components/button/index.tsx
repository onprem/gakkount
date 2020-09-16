import React from "react";

import styles from "./button.module.css";

export type ButtonProps = {
  variant?: "link" | "default";
} & React.DetailedHTMLProps<React.ButtonHTMLAttributes<HTMLButtonElement>, HTMLButtonElement>;

export const Button: React.FC<ButtonProps> = ({ children, variant = "default", ...rest }) => {
  return (
    <button {...rest} className={styles[variant]}>
      {children}
    </button>
  );
};

export default Button;
