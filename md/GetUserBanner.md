![image](https://github.com/bookWorm21/BannerService/assets/60971260/80b5059a-d56e-41b7-a529-8668902b744a)# [Главная](../README.md)

Для тестов были добавлены два баннера,
Banner: Id = 1, tags = [1, 2], feature = 1, enabled = true, content = '{"title": "some_title", "text": "some_text", "url": "some_url"}'
Banner: Id = 2, tags = [3], feature = 1, enabled = false, content = '{"text": "warning_banner"}'

1. Результат запроса - Unauthorized, так как не передан token авторизации:
![image](https://github.com/bookWorm21/BannerService/assets/60971260/55b9417a-121f-4593-b93c-61f282ae20cc)
2. Результат запроса - NoAccess, так как передан неверный token:
![image](https://github.com/bookWorm21/BannerService/assets/60971260/f2d9206b-a947-49c1-ab78-c3c50ef4e9a2)
3. При ошибках в параметрах:
![image](https://github.com/bookWorm21/BannerService/assets/60971260/98d8cd87-3efe-491a-bd55-9fff082d53e7)
![image](https://github.com/bookWorm21/BannerService/assets/60971260/11a1ec7c-9d62-4fe5-a275-d50660240095)
4. Контент баннера при успешном запросе:
![image](https://github.com/bookWorm21/BannerService/assets/60971260/27e3744e-ecfc-4d77-b3da-000c50e48070)
