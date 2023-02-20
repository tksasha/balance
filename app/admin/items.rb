# frozen_string_literal: true

ActiveAdmin.register Item do
  include HasCurrencyScopes

  decorate_with ItemDecorator

  menu priority: 2, label: proc { I18n.t('active_admin.items') }

  actions :all, except: %i[destroy new create]

  permit_params %i[date]

  includes :category

  filter :category, as: :select, collection: []

  index row_class: proc { 'item' } do
    column :date
    column :sum, class: 'sum' do |item|
      money(item.sum)
    end
    column :category, class: 'category'
    column :description do |item|
      sanitize(item.description)
    end
    actions
  end

  show title: :id do
    attributes_table do
      row :date
      row :sum do |item|
        money(item.sum)
      end
      row :category
      row(:description) { sanitize(item.description) }
      row :currency
      row :created_at
      row :updated_at
      row :deleted_at if resource.deleted_at.present?
    end
  end
end
