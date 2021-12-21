class CreateTags < ActiveRecord::Migration[7.0]
  def change
    create_table :tags do |t|
      t.references :category, foreign_key: true, null: false
      t.string :name, null: false

      t.timestamps
    end

    add_index :tags, [:category_id, :name], unique: true
  end
end
