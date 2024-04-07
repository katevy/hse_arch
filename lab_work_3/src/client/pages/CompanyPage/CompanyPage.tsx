import React from 'react';
import { CompanyTotalInformation } from '../../modules/CompanyTotalInformation';

const CompanyPage = () => {
    return (
        <div>
            <CompanyPageHeader />
            <CompanyTotalInformation />
            <CompanyPageFooter />
        </div>
    );
};

export default CompanyPage;