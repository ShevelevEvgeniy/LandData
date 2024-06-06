SELECT
    date_formation,
    (SELECT COUNT(*) FROM land_plots WHERE cad_number LIKE ($1 || '%')) AS amount_land_plots
FROM
    kpt
WHERE
        cad_quarter = $1;
