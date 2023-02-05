# frozen_string_literal: true

class ItemsController < BaseController
  def index
    respond_to :html, :js
  end

  def create
    respond_to :js

    render :new unless resource.save
  end

  def update
    respond_to :js

    render :edit unless resource.update(resource_params)
  end

  def destroy
    respond_to :js

    resource.destroy
  end

  private

  def scope
    Item
      .includes(:category)
      .order(date: :desc)
  end

  def collection
    @collection ||= ::ItemSearcher.search(scope, params)
  end

  def resource
    @resource ||= Item.find(params[:id])
  end

  def resource_params
    params
      .require(:item)
      .permit(:date, :formula, :category_id, :description, tag_ids: [])
      .merge(currency: params[:currency])
  end

  def build_resource
    @resource = Item.new(resource_params)
  end
end
