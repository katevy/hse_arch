import { createSlice } from '@reduxjs/toolkit';
import { useGetCompanyDataQuery } from './apiSlice';

interface CompanyTotalInformationState {
    data: any;
    error: string | null;
    loading: boolean;
}

const initialState: CompanyTotalInformationState = {
    data: null,
    error: null,
    loading: false,
};

const companyTotalInformationSlice = createSlice({
    name: 'companyTotalInformation',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
            .addMatcher(
                useGetCompanyDataQuery.matchPending,
                (state) => {
                    state.loading = true;
                    state.error = null;
                }
            )
            .addMatcher(
                useGetCompanyDataQuery.matchFulfilled,
                (state, action) => {
                    state.loading = false;
                    state.data = action.payload;
                }
            )
            .addMatcher(
                useGetCompanyDataQuery.matchRejected,
                (state, action) => {
                    state.loading = false;
                    state.error = action.error.message || 'Failed to fetch company data';
                }
            );
    },
});

export default companyTotalInformationSlice.reducer;
