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
}

const RunSettings = ({ tasks, indexes, onResume, onReset, isLoading, pointsStart, pointsEnd, pointsStep }: RunSettingsProps) => {
    const [selectedTasks, setSelectedTasks] = useState<string[]>([]);
    const [selectedIndexes, setSelectedIndexes] = useState<string[]>([]);

    // Set default selections when tasks or indexes change
    useEffect(() => {
        if (tasks.length > 0 && selectedTasks.length === 0) {
            setSelectedTasks(tasks.slice(0, 2).map(task => task.info.shortName));
        }
    }, [tasks]);

    useEffect(() => {
        if (indexes.length > 0 && selectedIndexes.length === 0) {
            setSelectedIndexes(indexes.slice(0, 2).map(index => index.info.shortName));
        }
    }, [indexes]);

    const handleTaskChange = (taskName: string) => {
        setSelectedTasks(prev =>
            prev.includes(taskName)
                ? prev.filter(t => t !== taskName)
                : [...prev, taskName]
        );
    };

    const handleIndexChange = (indexName: string) => {
        setSelectedIndexes(prev =>
            prev.includes(indexName)
                ? prev.filter(i => i !== indexName)
                : [...prev, indexName]
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

            <div className="button-group">
                <button
                    onClick={handleResume}
                    disabled={isLoading}
                    className="resume-button"
                >
                    ▶
                </button>
                <button
                    onClick={onReset}
                    disabled={isLoading}
                    className="reset-button"
                >
                    ⏹
                </button>
            </div>
        </div>
    );
};

export default RunSettings; 