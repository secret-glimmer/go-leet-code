import pandas as pd

def top_three_salaries(employee: pd.DataFrame, department: pd.DataFrame) -> pd.DataFrame:
    duplicate =employee.drop_duplicates(["salary","departmentId"])
    df_sorted = duplicate.sort_values(by=["salary"], ascending= False)
    df_groupby = df_sorted.groupby("departmentId").head(3)
    df = employee.merge(department, how ="left", left_on = "departmentId",right_on = "id")
    merging = df.merge(df_groupby[["salary","departmentId"]], how ="inner", on = ["departmentId","salary"])
    merging =  merging.rename(columns ={"name_y":"Department", "name_x": "Employee", "salary":"Salary"})
    return merging[["Department","Employee", "Salary"]]