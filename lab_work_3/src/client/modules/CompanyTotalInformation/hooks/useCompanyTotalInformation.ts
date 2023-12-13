import { useDispatch, useSelector } from 'react-redux';
import { RootState } from '../store';
import { addFavorite } from '../store/favoritesSlice';
import { useGetCompanyQuery } from '../services/api';

export const useCompanyTotalInformation = () => {
    const dispatch = useDispatch();
    const { data, error } = useGetCompanyQuery(companyName);
    const favorites = useSelector((state: RootState) => state.favorites);


    const handleAddToFavorite = () => {
        dispatch(addFavorite(data.company));
    };

    return { handleAddToFavorite, data, error }
}