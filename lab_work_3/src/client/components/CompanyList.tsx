import React from 'react';
import CompanyItem from '../ui/CompanyItem';

interface CompanyListProps {
    companies: string[];
    onCompanyClick: (companyName: string) => void;
}

const CompanyList: React.FC<CompanyListProps> = ({ companies, onCompanyClick }) => {
    return (
        <div>
            {companies.map((company) => (
                <CompanyItem key={company} companyName={company} onClick={() => onCompanyClick(company)} />
            ))}
        </div>
    );
};

export default CompanyList;