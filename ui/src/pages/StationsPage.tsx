import React, { useEffect, useState } from 'react';
import { Container, Typography, Grid, Card, CardContent, CardActions, Button } from '@mui/material';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

interface Station {
    id: number;
    name: string;
    tags: string[];
}

const StationsPage: React.FC = () => {
    const [stations, setStations] = useState<Station[]>([]);
    const navigate = useNavigate();

    useEffect(() => {
        // Fetch stations from the backend
        axios.get<Station[]>('http://localhost:8080/stations')
            .then(response => {
                setStations(response.data);
            })
            .catch(error => {
                console.error('There was an error fetching the stations!', error);
            });
    }, []);

    return (
        <Container>
            <Typography variant="h4" component="h1" gutterBottom>
                Stations
            </Typography>
            <Grid container spacing={4}>
                {stations.map((station) => (
                    <Grid item key={station.id} xs={12} sm={6} md={4}>
                        <Card>
                            <CardContent>
                                <Typography variant="h5" component="div">
                                    {station.name}
                                </Typography>
                                <Typography variant="body2" color="text.secondary">
                                    Tags: {station.tags.join(', ')}
                                </Typography>
                            </CardContent>
                            <CardActions>
                                <Button size="small" onClick={() => navigate(`/stations/${station.id}/songs`)}>View Songs</Button>                            </CardActions>
                        </Card>
                    </Grid>
                ))}
            </Grid>
        </Container>
    );

};

export default StationsPage;
