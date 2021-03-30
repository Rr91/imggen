 package config

type CinfI iterface {
	GetPost() stgin 
}

type Conf struct{
	port string
}

func New(confPath string){*Conf , error}{
	err := godotenv.Load(confPath)
	if err != nil{
		return nil, err
	}

	port := os.Getenv("SERVER_POST")

	if post == ""{
		return nil, err
	}
	return &Conf{port: port} , nil)
}