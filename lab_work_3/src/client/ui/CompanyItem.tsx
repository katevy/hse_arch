import React from 'react';

interface CompanyItemProps {
    companyName: string;
    onClick: () => void;
}

const CompanyItem: React.FC<CompanyItemProps> = ({ companyName, onClick }) => {
    return (
        <div onClick={onClick}>
            {companyName}
        </div>
    );
};

export default CompanyItem;