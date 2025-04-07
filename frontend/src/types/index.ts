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

export interface RunSettingsType {
    tasks: string[];
    indexes: string[];
    start: number;
    end: number;
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