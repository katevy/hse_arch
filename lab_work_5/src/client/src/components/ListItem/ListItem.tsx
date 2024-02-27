import './list-item.css'

interface ListItemProps {
    content: string
    id: number
    onDelete: (id: number) => void;
}

const ListItem = ({ content, id, onDelete }: ListItemProps) => {

    const handleOnClick = () => {
        onDelete(id)
    }

    return (
        <div className='list-item'>
            <span>{content}</span>
            <button onClick={handleOnClick} className='list-item_button'>Удалить</button>
        </div>
    );
};

export default ListItem;