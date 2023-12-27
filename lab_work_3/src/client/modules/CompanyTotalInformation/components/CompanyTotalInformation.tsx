import React from 'react';
import { useCompanyTotalInformation } from '../hooks/useCompanyTotalInformation';


interface CompanyTotalInformationProps {
    companyName: string;
}

const CompanyTotalInformation: React.FC<CompanyTotalInformationProps> = ({ companyName }) => {
    const { data, handleAddToFavorite, error } = useCompanyTotalInformation();

    if (error) {
        return <div>Error loading company information</div>;
    }

    if (!data) {
        return <div>Loading...</div>;
    }

    return (
        <div>
            <h2>{data.company.name}</h2>
            {/* Display other financial data as needed */}
            <button onClick={handleAddToFavorite}>Add to Favorites</button>
        </div>
    );
};
export default CompanyTotalInformation;
