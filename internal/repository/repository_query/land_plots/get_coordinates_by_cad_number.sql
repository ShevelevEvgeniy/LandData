SELECT ST_AsText(coordinates) FROM land_plots
WHERE cad_number = $1