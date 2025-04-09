import React, { useState } from 'react';
import { Run } from '../types';

interface RunsListProps {
    runs: Run[];
}

const RunsList: React.FC<RunsListProps> = ({ runs }) => {
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
                <div key={run.id} className="run-item">
                    <div
                        className="run-header"
                        onClick={() => toggleExpand(run.id)}
                    >
                        <span className="expand-icon">
                            {expandedIds.has(run.id) ? '▼' : '▶'}
                        </span>
                        <span className="run-id">#{run.id}</span>
                        <span className="run-status">{run.Status}</span>
                        <div className="run-preview">
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
                            Created at: {new Date(run.createdAt).toLocaleString()}
                        </div>
                    )}
                </div>
            ))}
        </div>
    );
};

export default RunsList; 