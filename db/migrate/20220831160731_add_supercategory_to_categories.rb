class AddSupercategoryToCategories < ActiveRecord::Migration[7.0]
  def change
    add_column :categories, :supercategory, :integer, null: false, default: 1
  end
end
