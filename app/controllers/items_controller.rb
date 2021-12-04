# frozen_string_literal: true

class ItemsController < ApplicationController
  delegate :destroy, to: :resource

  private

  def collection
    @collection ||= ::Items::GetCollectionService.call(params)
  end

  def result
    @result ||= ::Items::GetResultService.call(action_name, params)
  end
end
