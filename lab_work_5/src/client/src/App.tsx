import { useEffect, useState } from 'react';
import { SearchHistory, useDeleteSearchHistoryMutation, useGetSearchHistoriesQuery } from './api/searchHistory.api.slice';
import List from './components/List/List';

const App = () => {

  const [deleteSearchHistory, response] = useDeleteSearchHistoryMutation()
  const { data } = useGetSearchHistoriesQuery()

  const [items, setItems] = useState<SearchHistory[]>([])

  useEffect(() => {
    if (data)
      setItems(data)
  }, [data])

  useEffect(() => {
    if (response.isSuccess)
      alert('Данные удалены')
    else if (response.isError) {
      alert('Ошибка удаления')
    }
  }, [response.isError, response.isSuccess])

  const handleDeleteFunction = (id: number) => {
    // if (confirm('Удалить?'))
    deleteSearchHistory(id)
  }

  return (
    <div>
      <List items={items} deleteFunction={handleDeleteFunction} />
    </div>
  );
}

export default App;
