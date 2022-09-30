get-content .env | foreach {
    if(!$_.startswith("#")) {
        $name, $value = $_.split('=')
        set-content env:\$name $value
    }
}