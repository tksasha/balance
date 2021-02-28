class RemoveCategoriesSlug < ActiveRecord::Migration[6.1]
  def change
    if column_exists? :categories, :slug
      remove_column :categories, :slug
    end
  end
end
