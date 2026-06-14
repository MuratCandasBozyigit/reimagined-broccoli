function Food(){

    const food1 = "Tantuni"
    const food2 = "Kebap"

    return(
        <ul>
            <li>Ciğer</li>
            <li>{food1}</li>
            <li>{food2.toUpperCase()}</li>
        </ul>
    );
}

export default Food