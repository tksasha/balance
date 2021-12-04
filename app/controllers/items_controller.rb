# frozen_string_literal: true

class ItemsController < ApplicationController
  delegate :destroy, to: :resource

  def update
    render :edit, status: :unprocessable_entity unless resource.update resource_params
  end

  private

  def collection
    @collection ||= ::Items::GetCollectionService.call(params)
  end

  def result
    @result ||= ::Items::GetResultService.call(action_name, params)
  end
end
