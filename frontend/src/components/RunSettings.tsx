import { useState } from 'react';
import { Task, Index, RunSettings as RunSettingsType } from '../types';

interface RunSettingsProps {
    tasks: Task[];
    indexes: Index[];
    onResume: (settings: RunSettingsType) => void;
    isLoading: boolean;
}

const RunSettings = ({ tasks, indexes, onResume, isLoading }: RunSettingsProps) => {
    const [selectedTasks, setSelectedTasks] = useState<string[]>([]);
    const [selectedIndexes, setSelectedIndexes] = useState<string[]>([]);
    const [pointsStart, setPointsStart] = useState(1000);
    const [pointsEnd, setPointsEnd] = useState(10000);
    const [pointsStep, setPointsStep] = useState(100);

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
            end: pointsEnd,
            step: pointsStep
        });
    };

    return (
        <div className="run-settings">
            <div className="points-input">
                <div>
                    <label htmlFor="pointsStart">Начальное кол-во точек:</label>
                    <input
                        type="number"
                        id="pointsStart"
                        value={pointsStart}
                        onChange={(e) => setPointsStart(Number(e.target.value))}
                        min="100"
                        max="1000000"
                    />
                </div>
                <div>
                    <label htmlFor="pointsEnd">Конечное кол-во точек:</label>
                    <input
                        type="number"
                        id="pointsEnd"
                        value={pointsEnd}
                        onChange={(e) => setPointsEnd(Number(e.target.value))}
                        min="100"
                        max="1000000"
                    />
                </div>
                <div>
                    <label htmlFor="pointsStep">Шаг:</label>
                    <input
                        type="number"
                        id="pointsStep"
                        value={pointsStep}
                        onChange={(e) => setPointsStep(Number(e.target.value))}
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

            <button onClick={handleResume} disabled={isLoading}>
                ⏯
            </button>
        </div>
    );
};

export default RunSettings; 