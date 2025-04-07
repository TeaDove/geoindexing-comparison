import React, { useState, useEffect } from 'react';

export interface NotificationMessage {
    status: number;
    endpoint: string;
    method: string;
    error?: string;
    timestamp: number;
}

interface NotificationProps {
    message: NotificationMessage | null;
}

const Notification: React.FC<NotificationProps> = ({ message }) => {
    const [notifications, setNotifications] = useState<NotificationMessage[]>([]);

    useEffect(() => {
        if (message) {
            setNotifications(prev => [...prev, message]);
            const duration = message.status >= 400 ? 10000 : 2000;
            const timer = setTimeout(() => {
                setNotifications(prev => prev.filter(n => n.timestamp !== message.timestamp));
            }, duration);
            return () => clearTimeout(timer);
        }
    }, [message]);

    if (notifications.length === 0) return null;

    return (
        <div className="notification-container">
            {notifications.map((notification) => (
                <div
                    key={notification.timestamp}
                    className="notification"
                    style={{
                        borderColor: notification.status >= 500 ? '#dc3545' :
                            notification.status >= 400 ? '#ffc107' :
                                notification.status >= 200 ? '#28a745' : '#007bff'
                    }}
                >
                    <div className="notification-content">
                        <div className="notification-header">
                            <span className="status-dot" style={{
                                backgroundColor: notification.status >= 500 ? '#dc3545' :
                                    notification.status >= 400 ? '#ffc107' :
                                        notification.status >= 200 ? '#28a745' : '#007bff'
                            }} />
                            <span className="endpoint">{notification.method} {notification.endpoint}</span>
                            <span className="status">{notification.status}</span>
                        </div>
                        {notification.error && (
                            <div className="error-message">{notification.error}</div>
                        )}
                    </div>
                </div>
            ))}
        </div>
    );
};

export default Notification; 