# frozen_string_literal: true

RSpec.describe Category, type: :model do
  it { should validate_presence_of :name }

  it { should validate_uniqueness_of(:name).case_insensitive.scoped_to(:currency) }

  it { should define_enum_for(:currency).with_values(%w[uah usd rur]) }

  it { should callback(:assign_slug).before(:save) }

  describe '.visible' do
    let(:sql) { described_class.where(visible: true).to_sql }

    subject { described_class.visible.to_sql }

    it { should eq sql }
  end

  describe '.income' do
    let(:sql) { described_class.where(income: true).to_sql }

    subject { described_class.income.to_sql }

    it { should eq sql }
  end

  describe '.expense' do
    let(:sql) { described_class.where(income: false).to_sql }

    subject { described_class.expense.to_sql }

    it { should eq sql }
  end

  describe '#assign_slug' do
    subject { stub_model described_class, name: 'Інші надходження' }

    before { subject.send :assign_slug }

    its(:slug) { should eq 'inshi-nadkhodzhennia' }
  end
end
