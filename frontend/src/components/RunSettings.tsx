import { useState, useEffect } from 'react';
import type { Task, Index, RunSettings } from '../types/index';
import { formatNumber } from '../utils';

interface RunSettingsProps {
    tasks: Task[];
    indexes: Index[];
    onResume: (settings: RunSettings) => void;
    onReset: () => void;
    isLoading: boolean;
}

const RunSettings = ({ tasks, indexes, onResume, onReset, isLoading }: RunSettingsProps) => {
    const [selectedTasks, setSelectedTasks] = useState<string[]>([]);
    const [selectedIndexes, setSelectedIndexes] = useState<string[]>([]);
    const [pointsStart, setPointsStart] = useState(1000);
    const [pointsEnd, setPointsEnd] = useState(10000);
    const [pointsStep, setPointsStep] = useState(100);
    const [displayStart, setDisplayStart] = useState(formatNumber(1000));
    const [displayEnd, setDisplayEnd] = useState(formatNumber(10000));
    const [displayStep, setDisplayStep] = useState(formatNumber(100));

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

    const handleNumberInput = (value: string, setter: (n: number) => void, displaySetter: (s: string) => void) => {
        const numericValue = parseInt(value.replace(/\s/g, ''), 10);
        if (!isNaN(numericValue)) {
            setter(numericValue);
            displaySetter(formatNumber(numericValue));
        }
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
            <div className="points-input">
                <div>
                    <label htmlFor="pointsStart">Начальное кол-во точек</label>
                    <input
                        type="text"
                        id="pointsStart"
                        value={displayStart}
                        onChange={(e) => handleNumberInput(e.target.value, setPointsStart, setDisplayStart)}
                        min="100"
                        max="1000000"
                    />
                </div>
                <div>
                    <label htmlFor="pointsEnd">Конечное кол-во точек</label>
                    <input
                        type="text"
                        id="pointsEnd"
                        value={displayEnd}
                        onChange={(e) => handleNumberInput(e.target.value, setPointsEnd, setDisplayEnd)}
                        min="100"
                        max="1000000"
                    />
                </div>
                <div>
                    <label htmlFor="pointsStep">Шаг</label>
                    <input
                        type="text"
                        id="pointsStep"
                        value={displayStep}
                        onChange={(e) => handleNumberInput(e.target.value, setPointsStep, setDisplayStep)}
                        min="1"
                        max="1000"
                    />
                </div>
            </div>

            <fieldset id="indexes">
                <legend>Индексы</legend>
                {indexes.map(index => (
                    <div key={index.info.shortName}>
                        <input
                            type="checkbox"
                            id={`index-${index.info.shortName}`}
                            checked={selectedIndexes.includes(index.info.shortName)}
                            onChange={() => handleIndexChange(index.info.shortName)}
                        />
                        <label htmlFor={`index-${index.info.shortName}`}>
                            {index.info.longName}
                        </label>
                    </div>
                ))}
            </fieldset>

            <fieldset id="tasks">
                <legend>Задачи</legend>
                {tasks.map(task => (
                    <div key={task.info.shortName}>
                        <input
                            type="checkbox"
                            id={`task-${task.info.shortName}`}
                            checked={selectedTasks.includes(task.info.shortName)}
                            onChange={() => handleTaskChange(task.info.shortName)}
                        />
                        <label htmlFor={`task-${task.info.shortName}`}>
                            {task.info.longName}
                        </label>
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