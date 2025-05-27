import { useState, useEffect } from 'react';
import type { Task, Index, RunSettings } from '../types/index';
import './RunSettings.css';

interface RunSettingsProps {
    tasks: Task[];
    indexes: Index[];
    onResume: (settings: RunSettings) => void;
    onReset: () => void;
    isLoading: boolean;
    pointsStart: number;
    pointsEnd: number;
    pointsStep: number;
    setPointsStart: (n: number) => void;
    setPointsEnd: (n: number) => void;
    setPointsStep: (n: number) => void;
    selectedTasks: string[];
    selectedIndexes: string[];
    onTasksChange: (tasks: string[]) => void;
    onIndexesChange: (indexes: string[]) => void;
}

const RunSettings = ({
    tasks,
    indexes,
    onResume,
    onReset,
    isLoading,
    pointsStart,
    pointsEnd,
    pointsStep,
    selectedTasks,
    selectedIndexes,
    onTasksChange,
    onIndexesChange
}: RunSettingsProps) => {
    // Set default selections when tasks or indexes change
    useEffect(() => {
        if (tasks.length > 0 && selectedTasks.length === 0) {
            onTasksChange(tasks.slice(0, 2).map(task => task.info.shortName));
        }
    }, [tasks]);

    useEffect(() => {
        if (indexes.length > 0 && selectedIndexes.length === 0) {
            onIndexesChange(indexes.slice(0, 2).map(index => index.info.shortName));
        }
    }, [indexes]);

    const handleTaskChange = (taskName: string) => {
        onTasksChange(
            selectedTasks.includes(taskName)
                ? selectedTasks.filter(t => t !== taskName)
                : [...selectedTasks, taskName]
        );
    };

    const handleIndexChange = (indexName: string) => {
        onIndexesChange(
            selectedIndexes.includes(indexName)
                ? selectedIndexes.filter(i => i !== indexName)
                : [...selectedIndexes, indexName]
        );
    };

    const handleResume = () => {
        onResume({
            tasks: selectedTasks,
            indexes: selectedIndexes,
            start: pointsStart,
            stop: pointsEnd,
            step: pointsStep
        });
    };

    return (
        <div className="run-settings">
            <fieldset id="indexes">
                <legend>Индексы</legend>
                {indexes.map(index => (
                    <div key={index.info.shortName} className="tooltip-container">
                        <input
                            type="checkbox"
                            id={`index-${index.info.shortName}`}
                            checked={selectedIndexes.includes(index.info.shortName)}
                            onChange={() => handleIndexChange(index.info.shortName)}
                        />
                        <label htmlFor={`index-${index.info.shortName}`}>{index.info.longName}</label>
                        <span className="custom-tooltip">{index.info.description || 'Описание не доступно'}</span>
                    </div>
                ))}
            </fieldset>

            <fieldset id="tasks">
                <legend>Задачи</legend>
                {tasks.map(task => (
                    <div key={task.info.shortName} className="tooltip-container">
                        <input
                            type="checkbox"
                            id={`task-${task.info.shortName}`}
                            checked={selectedTasks.includes(task.info.shortName)}
                            onChange={() => handleTaskChange(task.info.shortName)}
                        />
                        <label htmlFor={`task-${task.info.shortName}`}>{task.info.longName}</label>
                        <span className="custom-tooltip">{task.info.description || 'Описание не доступно'}</span>
                    </div>
                ))}
            </fieldset>
        </div>
    );
};

export default RunSettings; 