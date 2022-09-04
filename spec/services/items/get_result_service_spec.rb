# frozen_string_literal: true

RSpec.describe Items::GetResultService do
  subject { described_class.new action_name, params }

  let(:params) { double }

  xdescribe '#call' do
    context 'when "create"' do
      let(:action_name) { 'create' }

      before { allow(Items::CreateService).to receive(:call).with(params).and_return(:result) }

      its(:call) { is_expected.to eq :result }
    end

    context 'when "edit"' do
      let(:action_name) { 'edit' }

      before { allow(Items::GetResourceService).to receive(:call).with(params).and_return(:result) }

      its(:call) { is_expected.to eq :result }
    end

    context 'when "update"' do
      let(:action_name) { 'update' }

      before { allow(Items::UpdateService).to receive(:call).with(params).and_return(:result) }

      its(:call) { is_expected.to eq :result }
    end

    context 'when "destroy"' do
      let(:action_name) { 'destroy' }

      before { allow(Items::DestroyService).to receive(:call).with(params).and_return(:result) }

      its(:call) { is_expected.to eq :result }
    end
  end
end
