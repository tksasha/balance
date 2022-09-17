# frozen_string_literal: true

module CategoriesHelper
  def supercategories
    Category
      .supercategories
      .each_with_object({}) do |i, res|
        name, = i

        res[I18n.t(name, scope: 'category.supercategory')] = name
      end
  end
end
