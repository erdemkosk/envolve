
![Logo](https://github.com/erdemkosk/envolve-ts/raw/master/logo.png)


# Envolve-Go

Synchronize environment variables across projects and **manage .env files**.
Automates the restoration of .env files for all projects, ensures consistency by syncing variables from a global .env file, and **creates symbolic links** to **maintain the latest environment settings**.


## Motivation


In today's fast-paced world, the demand for reliable and efficient software solutions continues to grow. **As developers, we strive not only to meet these demands** but to exceed them by creating robust applications that solve real-world problems.

1. **Disorganization**: Multiple .env files scattered across various project folders can lead to disorganization and confusion.

2. **Configuration Changes**: Over time, you may need to update environment variables across multiple projects. Manually making these changes in each .env file is tedious and error-prone.

3. **Data Loss**: When you delete or archive a project, you risk losing the associated .env files and their crucial configuration data.

The **motivation** behind our project stems from the desire to streamline and enhance the development process. We understand the **challenges developers face daily** – from **managing configurations** to ensuring seamless deployment. Our goal is to **simplify** these complexities and empower developers to focus more on innovation and less on repetitive tasks.

Join us on this journey of innovation and efficiency as we continue to push the boundaries of what's possible in software development.



## Envolve's Solution

**Envolve** aims to address these issues by providing a streamlined solution:

1. **Centralization**: Envolve centralizes all your .env files in a dedicated folder, making it easy to find and manage them.

2. **Symlink Support**: Envolve allows you to create symbolic links to your .env files, ensuring that you don't lose crucial configuration data when projects are deleted or archived.

3. **Visualization**: With Envolve, you can view the content of .env files in an organized tabular format for better clarity.

## Related

Here are some related projects

[Golang Version](https://github.com/erdemkosk/envolve-go)


## Roadmap

- Adding restore env files

- Adding drive upload function with enc

## Contributors

A big thank you to all the contributors who have helped make Envolve better:

<table>
  <tr>
    <td align="center">
      <a href="https://github.com/erdemkosk">
        <img src="https://github.com/erdemkosk.png" width="100px;" alt="erdemkosk"/>
        <br />
        <sub><b>Erdem Köşk</b></sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/suleymantaspinar">
        <img src="https://github.com/suleymantaspinar.png" width="100px;" alt="suleymantaspinar"/>
        <br />
        <sub><b>Süleyman Taşpınar</b></sub>
      </a>
    </td>
  </tr>
</table>
## Deployment

Release on Brew

```bash
  export GITHUB_TOKEN=xxxx
  goreleaser --snapshot  --clean  //it will create new snapshot
  git tag -a v1.0.0 -m "First release" && git push origin v1.0.0
  goreleaser release --clean 
```

