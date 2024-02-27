import { configureStore } from '@reduxjs/toolkit'
import { setupListeners } from '@reduxjs/toolkit/query'
import { searchHistoryApi } from '../api/searchHistory.api.slice'

export const store = configureStore({
    reducer: {
        [searchHistoryApi.reducerPath]: searchHistoryApi.reducer,
    },
    middleware: (getDefaultMiddleware) =>
        getDefaultMiddleware().concat(searchHistoryApi.middleware),
})

setupListeners(store.dispatch)