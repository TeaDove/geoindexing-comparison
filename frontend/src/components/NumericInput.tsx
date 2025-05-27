import React from 'react';
import './NumericInput.css';

interface NumericInputProps {
    id: string;
    value: number;
    onChange: (value: number) => void;
    min?: number;
    max?: number;
    label?: string;
    className?: string;
    disabled?: boolean;
}

const NumericInput: React.FC<NumericInputProps> = ({
    id,
    value,
    onChange,
    min,
    max,
    label,
    className,
    disabled
}) => {
    const formatNumber = (num: number): string => {
        return num.toString().replace(/\B(?=(\d{3})+(?!\d))/g, " ");
    };

    const parseNumber = (str: string): number => {
        return Number(str.replace(/\s/g, ''));
    };

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const rawValue = e.target.value;
        const parsedValue = parseNumber(rawValue);

        if (isNaN(parsedValue)) return;

        if (min !== undefined && parsedValue < min) return;
        if (max !== undefined && parsedValue > max) return;

        onChange(parsedValue);
    };

    return (
        <div className={`numeric-input ${className || ''}`}>
            {label && <label htmlFor={id}>{label}</label>}
            <input
                id={id}
                type="text"
                inputMode="numeric"
                value={formatNumber(value)}
                onChange={handleChange}
                disabled={disabled}
            />
        </div>
    );
};

export default NumericInput; 