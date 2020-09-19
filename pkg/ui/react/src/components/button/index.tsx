import React from "react";

import styles from "./button.module.css";

export type ButtonProps = {
  variant?: "link" | "pill" | "default";
} & React.DetailedHTMLProps<React.ButtonHTMLAttributes<HTMLButtonElement>, HTMLButtonElement>;

export const Button: React.FC<ButtonProps> = ({ children, variant = "default", className, ...rest }) => {
  return (
    <button {...rest} className={`${styles[variant]} ${className}`}>
      {children}
    </button>
  );
};

export default Button;
