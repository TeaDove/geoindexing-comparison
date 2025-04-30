import React, { useState } from 'react';
import { Run } from '../types';
import Charts from './Charts';

interface RunsListProps {
    runs: Run[];
    onRunSelect: (runId: number) => void;
    selectedRunId: number | null;
}

const RunsList: React.FC<RunsListProps> = ({ runs, onRunSelect, selectedRunId }) => {
    const [expandedIds, setExpandedIds] = useState<Set<number>>(new Set());

    const toggleExpand = (id: number) => {
        setExpandedIds(prev => {
            const newSet = new Set(prev);
            if (newSet.has(id)) {
                newSet.delete(id);
            } else {
                newSet.add(id);
            }
            return newSet;
        });
    };

    return (
        <div className="runs-list">
            {runs.map(run => (
                <div
                    key={run.id}
                    className={`run-item ${selectedRunId === run.id ? 'selected' : ''}`}
                >
                    <div
                        className="run-header"
                        onClick={() => {
                            toggleExpand(run.id);
                            onRunSelect(run.id);
                        }}
                    >
                        <span className="expand-icon">
                            {expandedIds.has(run.id) ? '▼' : '▶'}
                        </span>
                        <span className="run-id">#{run.id}</span>
                        <div className="run-preview">
                            <div className="run-status-preview">
                                <span className="label">Status:</span>
                                {run.Status}
                            </div>
                            {run.indexes.length > 0 && (
                                <div className="run-indexes">
                                    <span className="label">Indexes:</span>
                                    {run.indexes.join(', ')}
                                </div>
                            )}
                            {run.tasks.length > 0 && (
                                <div className="run-tasks">
                                    <span className="label">Tasks:</span>
                                    {run.tasks.join(', ')}
                                </div>
                            )}
                        </div>
                    </div>
                    {expandedIds.has(run.id) && (
                        <div className="run-details">
                            <div>Created at: {new Date(run.createdAt).toLocaleString()}</div>
                            {run.completedAt && (
                                <div>
                                    Completed at: {new Date(run.completedAt).toLocaleString()}
                                    {' '}
                                    ({((new Date(run.completedAt).getTime() - new Date(run.createdAt).getTime()) / 1000).toFixed(1)}s)
                                </div>
                            )}
                            <div>Points range: {run.start} - {run.stop}</div>
                            <div>Step: {run.step}</div>
                            <Charts selectedRunId={run.id} run={run} />
                        </div>
                    )}
                </div>
            ))}
        </div>
    );
};

export default RunsList; 