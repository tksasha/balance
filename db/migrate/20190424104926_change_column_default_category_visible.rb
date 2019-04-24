class ChangeColumnDefaultCategoryVisible < ActiveRecord::Migration[5.2]
  def change
    change_column_default :categories, :visible, 1
  end
end
