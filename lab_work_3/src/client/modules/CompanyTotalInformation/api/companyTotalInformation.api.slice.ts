import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

const BASE_URL = 'https://baseURL';

interface CompanyData {
    // структура данных
}

export const api = createApi({
    baseQuery: fetchBaseQuery({ baseUrl: BASE_URL }),
    endpoints: (builder) => ({
        getCompanyData: builder.query<CompanyData, string>({
            query: (companyName) => `company/${companyName}`, // Путь к вашему эндпоинту
        }),
    }),
});

export const { useGetCompanyDataQuery } = api;
