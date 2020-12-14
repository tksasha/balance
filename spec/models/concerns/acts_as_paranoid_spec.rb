# frozen_string_literal: true

RSpec.describe ActsAsParanoid, type: :model do
  let(:described_class) do
    Class.new(ApplicationRecord) do
      include ActsAsParanoid

      self.table_name = :items
    end
  end

  describe '#destroy' do
    before { expect(subject).to receive(:touch).with(:deleted_at) }

    it { expect { subject.destroy }.not_to raise_error }
  end

  describe '.default_scope' do
    let(:sql) { described_class.unscoped.where(deleted_at: nil).to_sql }

    subject { described_class.all }

    its(:to_sql) { should eq sql }
  end
end
