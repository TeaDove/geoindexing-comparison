export interface TaskInfo {
    shortName: string;
    longName: string;
}

export interface IndexInfo {
    shortName: string;
    longName: string;
}

export interface Task {
    info: TaskInfo;
}

export interface Index {
    info: IndexInfo;
}

export interface RunSettings {
    tasks: string[];
    indexes: string[];
    start: number;
    stop: number;
    step: number;
}

export interface ChartData {
    labels: string[];
    datasets: {
        label: string;
        data: number[];
        borderColor: string;
        backgroundColor: string;
    }[];
}

export interface Run {
    id: number;
    Status: string;
    indexes: string[];
    tasks: string[];
    createdAt: string;
    completedAt?: string;
    createdBy: string;
    start: number;
    stop: number;
    step: number;
}

// Added Point type based on Charts.tsx usage
export interface Point {
    task: string;
    index: string;
    x: number;
    y: number;
} 