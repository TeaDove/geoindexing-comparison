export interface Run {
    id: number;
    createdAt: string;
    createdBy: string;
    completedAt: string;
    Status: string;
    indexes: string[];
    tasks: string[];
    start: number;
    stop: number;
    step: number;
} 