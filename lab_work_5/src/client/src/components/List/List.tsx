import ListItem from "../ListItem/ListItem";

interface ListProps {
    items: any[]
    deleteFunction: (id: number) => void
}

const List = ({ items, deleteFunction }: ListProps) => {

    const handleDelete = (id: number) => {
        deleteFunction(id)
    }

    return (
        <div>
            {items.map((item) => (
                <ListItem id={item.ID} content={item.Content} onDelete={handleDelete} />
            ))}
        </div>
    );
};

export default List;